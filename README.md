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

最后更新：2026-03-07 10:26:19 UTC（2026-03-07 18:26:19 UTC+8）

**代理总数：71**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 121ms | ✓ 923ms | ✓ 1101ms | ✓ 1301ms | ✓ 809ms | http |
| 120.92.212.16:7890 | ✓ 1106ms | ✓ 1448ms | ✓ 1166ms | 否 | ✓ 1155ms | http |
| 61.72.221.234:3128 | ✓ 1407ms | 否 | ✓ 1724ms | 否 | ✓ 1057ms | http |
| 178.236.245.59:3128 | ✓ 1040ms | 否 | ✓ 1578ms | 否 | ✓ 1992ms | http |
| 91.107.175.112:10801 | ✓ 421ms | 否 | ✓ 1727ms | ✓ 1523ms | 否 | http |
| 168.235.110.63:3128 | 否 | ✓ 927ms | ✓ 1785ms | ✓ 935ms | ✓ 1728ms | http |
| 120.240.35.161:22222 | 否 | 否 | ✓ 1257ms | ✓ 1382ms | ✓ 1135ms | http |
| 167.172.69.123:8080 | ✓ 1593ms | 否 | 否 | ✓ 1236ms | ✓ 1023ms | http |
| 167.172.69.123:80 | ✓ 1667ms | 否 | 否 | ✓ 1354ms | ✓ 1001ms | http |
| 81.70.169.194:80 | 否 | ✓ 1526ms | ✓ 1200ms | ✓ 1423ms | ✓ 1173ms | http |
| 150.107.140.238:3128 | ✓ 1852ms | 否 | 否 | ✓ 1518ms | ✓ 1132ms | http |
| 185.115.74.185:8080 | ✓ 1064ms | ✓ 1795ms | ✓ 1269ms | 否 | 否 | http |
| 120.240.29.53:22222 | ✓ 1141ms | ✓ 1514ms | ✓ 1151ms | ✓ 1429ms | ✓ 1060ms | http |
| 94.72.109.169:8080 | ✓ 1040ms | 否 | 否 | ✓ 1501ms | ✓ 1418ms | http |
| 125.128.12.144:3128 | ✓ 1756ms | 否 | ✓ 1028ms | ✓ 1253ms | ✓ 994ms | http |
| 178.236.245.17:3128 | ✓ 1062ms | 否 | ✓ 1525ms | 否 | ✓ 1542ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1339ms | ✓ 1164ms | ✓ 1197ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1151ms | 否 | ✓ 1473ms | ✓ 1416ms | ✓ 1376ms | http |
| 34.101.184.164:3128 | ✓ 1759ms | 否 | ✓ 1764ms | 否 | ✓ 1191ms | http |
| 5.252.33.13:2025 | 否 | 否 | ✓ 1329ms | ✓ 1940ms | ✓ 1605ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1385ms | ✓ 1835ms | 否 | ✓ 1118ms | http |
| 190.9.109.199:999 | ✓ 1965ms | ✓ 1475ms | ✓ 1113ms | ✓ 1219ms | ✓ 1021ms | http |
| 116.80.80.197:3172 | ✓ 1732ms | 否 | ✓ 1686ms | 否 | ✓ 1817ms | http |
| 190.9.109.205:999 | ✓ 1967ms | 否 | ✓ 1114ms | ✓ 1157ms | 否 | http |
| 35.225.22.61:80 | ✓ 301ms | ✓ 1217ms | ✓ 236ms | ✓ 1071ms | ✓ 872ms | http |
| 120.232.242.119:22222 | ✓ 1061ms | ✓ 1397ms | 否 | ✓ 1362ms | 否 | http |
| 117.159.239.44:22222 | ✓ 1118ms | ✓ 1265ms | ✓ 984ms | ✓ 1479ms | ✓ 1033ms | http |
| 162.248.165.72:1080 | ✓ 496ms | 否 | ✓ 733ms | 否 | ✓ 1336ms | http |
| 88.80.150.82:8080 | ✓ 1460ms | ✓ 1887ms | 否 | ✓ 1952ms | ✓ 1592ms | https |
| 183.249.5.214:22222 | ✓ 1168ms | ✓ 1340ms | ✓ 1013ms | 否 | ✓ 1095ms | http |
| 91.233.223.147:3128 | ✓ 738ms | 否 | ✓ 691ms | ✓ 1869ms | ✓ 1430ms | http |
| 91.193.240.157:9877 | ✓ 1017ms | 否 | ✓ 917ms | 否 | ✓ 1451ms | http |
| 120.240.35.160:22222 | ✓ 1434ms | ✓ 1408ms | ✓ 1194ms | ✓ 1389ms | ✓ 1079ms | http |
| 103.215.36.88:19205 | ✓ 1435ms | 否 | 否 | ✓ 1615ms | ✓ 1218ms | http |
| 101.43.255.96:80 | ✓ 1134ms | 否 | ✓ 1826ms | ✓ 1898ms | ✓ 1526ms | http |
| 185.243.218.43:49153 | ✓ 857ms | 否 | ✓ 1923ms | ✓ 1905ms | 否 | http |
| 59.46.216.131:30001 | 否 | ✓ 1731ms | ✓ 1392ms | ✓ 1594ms | ✓ 1675ms | http |
| 104.248.243.244:3128 | ✓ 555ms | ✓ 1501ms | ✓ 718ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 952ms | ✓ 1672ms | ✓ 679ms | ✓ 1563ms | ✓ 1071ms | http |
| 2.63.162.206:3128 | ✓ 1254ms | 否 | ✓ 1869ms | ✓ 1947ms | 否 | http |
| 136.49.39.94:8888 | 否 | 否 | ✓ 1526ms | ✓ 1524ms | ✓ 1284ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1860ms | 否 | ✓ 1062ms | ✓ 828ms | http |
| 120.240.29.168:22222 | ✓ 1215ms | ✓ 1477ms | ✓ 1165ms | ✓ 1347ms | ✓ 1078ms | http |
| 212.175.29.184:8080 | ✓ 738ms | ✓ 1951ms | ✓ 1866ms | 否 | ✓ 1677ms | http |
| 120.240.35.177:22222 | ✓ 1081ms | ✓ 1416ms | ✓ 1036ms | ✓ 1708ms | 否 | http |
| 120.198.141.79:22222 | ✓ 1267ms | ✓ 1759ms | ✓ 1433ms | ✓ 1594ms | 否 | http |
| 175.0.74.111:10808 | ✓ 1088ms | 否 | ✓ 1561ms | 否 | ✓ 1913ms | http |
| 120.240.35.173:22222 | ✓ 1155ms | ✓ 1405ms | ✓ 1173ms | 否 | ✓ 1125ms | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 1309ms | ✓ 1230ms | ✓ 1008ms | http |
| 113.59.32.162:22222 | ✓ 1338ms | ✓ 1843ms | ✓ 1658ms | 否 | 否 | http |
| 188.191.147.116:8082 | ✓ 1801ms | ✓ 1646ms | ✓ 1956ms | 否 | 否 | http |
| 3.99.169.21:22841 | ✓ 1712ms | 否 | ✓ 1646ms | 否 | ✓ 1976ms | http |
| 85.9.195.140:1080 | ✓ 925ms | 否 | ✓ 893ms | 否 | ✓ 1748ms | http |
| 14.56.107.244:3128 | ✓ 960ms | 否 | ✓ 1550ms | ✓ 1320ms | 否 | http |
| 116.80.80.201:3172 | ✓ 1645ms | 否 | ✓ 1641ms | 否 | ✓ 1816ms | http |
| 117.159.239.51:22222 | ✓ 941ms | ✓ 1198ms | 否 | ✓ 1291ms | 否 | http |
| 107.172.125.217:3128 | ✓ 408ms | 否 | ✓ 1860ms | ✓ 958ms | ✓ 779ms | http |
| 46.183.25.8:443 | 否 | 否 | ✓ 629ms | ✓ 1096ms | ✓ 1369ms | http |
| 172.212.68.37:3128 | ✓ 375ms | 否 | ✓ 613ms | ✓ 1726ms | ✓ 1531ms | http |
| 222.184.48.252:22222 | ✓ 1137ms | 否 | ✓ 1022ms | ✓ 1674ms | 否 | http |
| 125.128.12.14:3128 | ✓ 1103ms | 否 | 否 | ✓ 1983ms | ✓ 1320ms | http |
| 61.72.221.94:3128 | ✓ 1073ms | ✓ 1434ms | ✓ 1865ms | 否 | 否 | http |
| 61.72.221.194:3128 | ✓ 1693ms | 否 | ✓ 1473ms | 否 | ✓ 992ms | http |
| 121.128.121.54:3128 | ✓ 1689ms | 否 | ✓ 1409ms | ✓ 1203ms | ✓ 965ms | http |
| 61.72.110.54:3128 | 否 | ✓ 1770ms | ✓ 1875ms | ✓ 1718ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1313ms | 否 | ✓ 602ms | ✓ 1420ms | ✓ 1154ms | http |
| 116.80.82.227:3172 | ✓ 1695ms | 否 | ✓ 1878ms | ✓ 1983ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1302ms | ✓ 1903ms | ✓ 1626ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1972ms | 否 | ✓ 1796ms | ✓ 1411ms | ✓ 1729ms | http |
| 120.240.29.173:22222 | ✓ 1132ms | ✓ 1386ms | ✓ 1104ms | ✓ 1342ms | ✓ 1152ms | http |
| 117.159.239.52:22222 | ✓ 1043ms | ✓ 1227ms | 否 | ✓ 1672ms | ✓ 1091ms | http |

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
