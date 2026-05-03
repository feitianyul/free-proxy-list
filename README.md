**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-05-03 11:00:38 UTC（2026-05-03 19:00:38 UTC+8）

**代理总数：72**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 434ms | 否 | ✓ 212ms | ✓ 956ms | 否 | http |
| 218.108.131.186:17890 | ✓ 999ms | ✓ 1272ms | ✓ 996ms | ✓ 1288ms | ✓ 1036ms | http |
| 206.206.126.177:2412 | ✓ 1107ms | ✓ 1787ms | ✓ 928ms | ✓ 1200ms | ✓ 962ms | http |
| 1.231.81.166:3128 | ✓ 1845ms | ✓ 1409ms | ✓ 1511ms | ✓ 1227ms | ✓ 1302ms | http |
| 212.58.132.5:8888 | ✓ 1202ms | 否 | ✓ 1373ms | ✓ 1483ms | ✓ 1254ms | http |
| 113.160.132.26:8080 | ✓ 1960ms | ✓ 1583ms | ✓ 1809ms | ✓ 1483ms | ✓ 1246ms | http |
| 45.167.124.71:999 | ✓ 1091ms | ✓ 1512ms | ✓ 1342ms | ✓ 1906ms | ✓ 1549ms | http |
| 47.77.216.82:1080 | ✓ 290ms | ✓ 1799ms | ✓ 1113ms | ✓ 1075ms | 否 | http |
| 109.120.156.122:8090 | ✓ 849ms | 否 | ✓ 1311ms | 否 | ✓ 1739ms | http |
| 45.153.231.229:8080 | ✓ 1006ms | 否 | ✓ 1636ms | 否 | ✓ 1921ms | http |
| 193.123.250.39:1080 | ✓ 1262ms | 否 | 否 | ✓ 1424ms | ✓ 1262ms | http |
| 46.105.190.38:3128 | ✓ 1982ms | ✓ 1757ms | ✓ 723ms | ✓ 1769ms | ✓ 1287ms | http |
| 217.76.245.80:999 | ✓ 733ms | ✓ 1573ms | ✓ 1064ms | ✓ 1876ms | ✓ 1242ms | http |
| 148.230.4.241:999 | ✓ 1626ms | 否 | ✓ 755ms | ✓ 1541ms | ✓ 1463ms | http |
| 59.46.216.131:30001 | ✓ 1164ms | ✓ 1621ms | ✓ 1400ms | ✓ 1880ms | ✓ 1226ms | http |
| 106.10.55.212:1121 | ✓ 1652ms | ✓ 1916ms | 否 | 否 | ✓ 1065ms | http |
| 94.131.118.39:1081 | ✓ 856ms | ✓ 1057ms | ✓ 1007ms | ✓ 1513ms | ✓ 1124ms | http |
| 38.180.62.47:10808 | ✓ 1897ms | 否 | ✓ 1835ms | 否 | ✓ 1095ms | http |
| 62.60.231.71:56608 | ✓ 1355ms | ✓ 1850ms | ✓ 1781ms | 否 | ✓ 1367ms | http |
| 121.230.9.225:1080 | ✓ 1522ms | ✓ 1474ms | ✓ 1137ms | ✓ 1551ms | ✓ 1138ms | http |
| 34.96.238.40:8080 | ✓ 1039ms | ✓ 1300ms | 否 | ✓ 1221ms | ✓ 1236ms | http |
| 8.219.97.248:80 | ✓ 1503ms | 否 | ✓ 1994ms | ✓ 1649ms | 否 | http |
| 77.110.107.80:8080 | ✓ 702ms | ✓ 1578ms | ✓ 1577ms | 否 | 否 | http |
| 77.110.107.80:1080 | ✓ 713ms | ✓ 1544ms | ✓ 1613ms | 否 | 否 | http |
| 147.45.178.211:14658 | 否 | 否 | ✓ 496ms | ✓ 1413ms | ✓ 1376ms | http |
| 62.133.60.126:24558 | ✓ 893ms | ✓ 1604ms | ✓ 1550ms | ✓ 1412ms | ✓ 1347ms | http |
| 62.60.149.161:3128 | ✓ 1428ms | ✓ 1976ms | ✓ 717ms | ✓ 1948ms | 否 | http |
| 173.212.246.157:3128 | ✓ 1801ms | 否 | ✓ 998ms | 否 | ✓ 1956ms | http |
| 38.180.121.135:10808 | ✓ 716ms | ✓ 1183ms | 否 | 否 | ✓ 1336ms | http |
| 152.70.91.193:40000 | ✓ 1288ms | 否 | ✓ 1810ms | 否 | ✓ 1485ms | http |
| 8.154.21.175:3128 | ✓ 1032ms | ✓ 1282ms | ✓ 1066ms | ✓ 1335ms | ✓ 1081ms | http |
| 45.125.67.37:8443 | ✓ 1064ms | 否 | ✓ 1267ms | ✓ 1221ms | ✓ 1247ms | http |
| 101.32.244.83:8080 | ✓ 1259ms | ✓ 1618ms | ✓ 1136ms | ✓ 1695ms | ✓ 1515ms | http |
| 121.43.196.213:8222 | ✓ 1107ms | ✓ 1232ms | ✓ 986ms | ✓ 1317ms | ✓ 1153ms | http |
| 121.43.196.210:8222 | ✓ 1094ms | ✓ 1311ms | ✓ 994ms | ✓ 1310ms | ✓ 1122ms | http |
| 121.230.8.22:1080 | ✓ 1335ms | ✓ 1491ms | ✓ 1130ms | ✓ 1533ms | ✓ 1382ms | http |
| 46.105.190.40:3128 | ✓ 458ms | ✓ 1251ms | ✓ 658ms | ✓ 1642ms | ✓ 1321ms | http |
| 154.64.232.35:8080 | ✓ 856ms | 否 | ✓ 1206ms | 否 | ✓ 736ms | http |
| 94.131.118.129:1082 | ✓ 1294ms | ✓ 1515ms | 否 | ✓ 1705ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1031ms | 否 | ✓ 1159ms | 否 | ✓ 1607ms | http |
| 154.90.48.209:9090 | 否 | 否 | ✓ 1514ms | ✓ 1726ms | ✓ 1160ms | http |
| 20.127.128.70:8080 | ✓ 1170ms | 否 | ✓ 1797ms | 否 | ✓ 1795ms | http |
| 105.159.137.18:4156 | ✓ 863ms | ✓ 1567ms | ✓ 1040ms | 否 | ✓ 1603ms | http |
| 105.159.137.18:5062 | ✓ 860ms | ✓ 1733ms | ✓ 873ms | 否 | ✓ 1646ms | http |
| 130.61.174.200:1080 | 否 | ✓ 1286ms | ✓ 1721ms | ✓ 1186ms | 否 | http |
| 20.164.75.153:8080 | ✓ 1684ms | 否 | ✓ 1896ms | 否 | ✓ 1998ms | http |
| 3.101.133.120:80 | ✓ 359ms | 否 | ✓ 526ms | ✓ 984ms | ✓ 784ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 921ms | ✓ 1497ms | ✓ 883ms | http |
| 45.140.147.155:1081 | ✓ 697ms | ✓ 1507ms | ✓ 1257ms | ✓ 1933ms | ✓ 1332ms | http |
| 37.187.109.70:10111 | ✓ 1359ms | 否 | ✓ 1692ms | 否 | ✓ 1650ms | http |
| 80.92.204.47:1081 | ✓ 390ms | 否 | ✓ 894ms | 否 | ✓ 1172ms | http |
| 202.141.161.53:10808 | ✓ 1115ms | ✓ 1468ms | ✓ 1673ms | ✓ 1385ms | 否 | http |
| 62.113.119.14:8080 | ✓ 565ms | ✓ 1341ms | ✓ 902ms | ✓ 1476ms | ✓ 1043ms | http |
| 138.124.99.216:8888 | 否 | ✓ 1894ms | ✓ 708ms | 否 | ✓ 1527ms | http |
| 178.156.224.42:3128 | ✓ 1125ms | ✓ 1583ms | ✓ 1463ms | ✓ 1801ms | ✓ 1552ms | http |
| 43.133.44.89:8888 | ✓ 1310ms | 否 | ✓ 1333ms | 否 | ✓ 1067ms | http |
| 34.101.184.164:3128 | ✓ 1020ms | 否 | ✓ 1056ms | ✓ 1530ms | ✓ 1313ms | http |
| 213.3.34.39:443 | ✓ 1167ms | 否 | ✓ 762ms | ✓ 1520ms | ✓ 1269ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1753ms | ✓ 1637ms | ✓ 1166ms | 否 | http |
| 118.31.1.154:80 | ✓ 1047ms | ✓ 1342ms | ✓ 1205ms | ✓ 1374ms | ✓ 1077ms | http |
| 135.125.97.184:35749 | ✓ 1638ms | ✓ 1786ms | ✓ 1683ms | 否 | 否 | http |
| 121.230.8.213:1080 | ✓ 1353ms | 否 | ✓ 1111ms | ✓ 1578ms | ✓ 1344ms | http |
| 121.230.8.97:1080 | ✓ 1155ms | ✓ 1815ms | ✓ 1334ms | ✓ 1575ms | ✓ 1320ms | http |
| 101.32.243.189:80 | ✓ 1635ms | 否 | ✓ 1512ms | ✓ 1455ms | ✓ 1487ms | http |
| 152.42.177.32:8888 | ✓ 1611ms | 否 | ✓ 1415ms | ✓ 1519ms | ✓ 1543ms | http |
| 121.230.8.144:1080 | 否 | ✓ 1613ms | ✓ 1852ms | ✓ 1496ms | ✓ 1346ms | http |
| 117.236.124.166:3128 | ✓ 1611ms | 否 | ✓ 1644ms | 否 | ✓ 1807ms | http |
| 105.159.154.44:4308 | ✓ 1656ms | ✓ 1476ms | ✓ 1693ms | 否 | ✓ 1554ms | http |
| 160.19.110.130:8082 | ✓ 1931ms | 否 | ✓ 1479ms | 否 | ✓ 1751ms | http |
| 217.182.195.221:30003 | ✓ 1224ms | 否 | ✓ 973ms | ✓ 1955ms | 否 | http |
| 45.129.141.143:3128 | ✓ 1716ms | 否 | ✓ 1249ms | 否 | ✓ 1779ms | http |
| 168.110.52.228:3128 | ✓ 1778ms | 否 | 否 | ✓ 1067ms | ✓ 930ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
