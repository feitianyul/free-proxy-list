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

最后更新：2026-03-06 19:44:16 UTC（2026-03-07 03:44:16 UTC+8）

**代理总数：79**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 834ms | ✓ 1118ms | ✓ 888ms | ✓ 1309ms | ✓ 987ms | http |
| 152.42.195.165:8888 | ✓ 1436ms | 否 | ✓ 809ms | ✓ 1103ms | ✓ 859ms | http |
| 194.213.18.200:443 | ✓ 1190ms | ✓ 1814ms | 否 | 否 | ✓ 1936ms | http |
| 91.233.223.147:3128 | ✓ 1044ms | 否 | ✓ 1936ms | 否 | ✓ 1706ms | http |
| 1.231.81.166:3128 | ✓ 852ms | ✓ 920ms | ✓ 827ms | ✓ 1063ms | ✓ 735ms | http |
| 192.166.82.55:1080 | ✓ 773ms | 否 | ✓ 1315ms | ✓ 1704ms | ✓ 1324ms | http |
| 186.148.180.46:999 | 否 | 否 | ✓ 1428ms | ✓ 1922ms | ✓ 1696ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 825ms | ✓ 1055ms | ✓ 1173ms | http |
| 85.9.195.140:1080 | ✓ 1297ms | 否 | ✓ 526ms | ✓ 1323ms | ✓ 1089ms | http |
| 168.235.110.63:3128 | ✓ 505ms | 否 | ✓ 1317ms | 否 | ✓ 1331ms | http |
| 125.128.12.144:3128 | ✓ 1495ms | ✓ 1388ms | ✓ 901ms | 否 | 否 | http |
| 165.227.5.10:8888 | ✓ 508ms | 否 | ✓ 577ms | 否 | ✓ 836ms | http |
| 46.183.25.8:443 | ✓ 1465ms | 否 | ✓ 1828ms | 否 | ✓ 1462ms | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 1379ms | ✓ 1257ms | ✓ 817ms | http |
| 167.172.69.123:8080 | ✓ 1211ms | 否 | ✓ 1095ms | ✓ 1158ms | 否 | http |
| 14.56.107.244:3128 | ✓ 1893ms | 否 | ✓ 700ms | ✓ 1032ms | ✓ 1795ms | http |
| 181.78.79.155:999 | ✓ 1088ms | ✓ 1788ms | ✓ 679ms | ✓ 1737ms | ✓ 1887ms | http |
| 96.9.66.244:8080 | ✓ 1409ms | 否 | ✓ 1547ms | ✓ 1861ms | ✓ 1723ms | http |
| 190.12.150.244:999 | ✓ 1408ms | ✓ 1891ms | ✓ 1873ms | ✓ 1689ms | ✓ 1754ms | http |
| 136.49.39.94:8888 | ✓ 1665ms | 否 | ✓ 692ms | ✓ 1276ms | ✓ 1224ms | http |
| 178.236.245.17:3128 | ✓ 1088ms | 否 | ✓ 1615ms | 否 | ✓ 1667ms | http |
| 178.236.245.59:3128 | ✓ 1088ms | 否 | ✓ 1494ms | 否 | ✓ 1748ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1101ms | ✓ 1316ms | ✓ 1031ms | ✓ 1241ms | http |
| 61.72.221.94:3128 | ✓ 1564ms | ✓ 1924ms | ✓ 1930ms | ✓ 1704ms | 否 | http |
| 121.128.121.54:3128 | ✓ 1556ms | 否 | ✓ 1030ms | ✓ 1404ms | 否 | http |
| 81.70.169.194:80 | ✓ 1056ms | ✓ 1297ms | ✓ 967ms | ✓ 1236ms | ✓ 987ms | http |
| 101.43.255.96:80 | ✓ 1187ms | ✓ 1461ms | ✓ 933ms | ✓ 1311ms | ✓ 971ms | http |
| 167.172.69.123:80 | ✓ 770ms | 否 | ✓ 1450ms | ✓ 1081ms | ✓ 1147ms | http |
| 61.72.110.94:3128 | ✓ 1586ms | 否 | ✓ 668ms | 否 | ✓ 1395ms | http |
| 94.176.3.43:7443 | ✓ 1031ms | 否 | ✓ 1671ms | ✓ 1941ms | 否 | http |
| 61.72.221.234:3128 | ✓ 1588ms | 否 | ✓ 1279ms | ✓ 1722ms | 否 | http |
| 103.104.99.29:80 | ✓ 1985ms | 否 | 否 | ✓ 1492ms | ✓ 1465ms | http |
| 125.128.12.14:3128 | ✓ 1268ms | ✓ 1341ms | 否 | 否 | ✓ 1044ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1656ms | ✓ 1127ms | ✓ 1422ms | http |
| 106.14.203.63:3333 | 否 | 否 | ✓ 888ms | ✓ 1115ms | ✓ 862ms | http |
| 61.72.110.54:3128 | 否 | 否 | ✓ 895ms | ✓ 1064ms | ✓ 1357ms | http |
| 91.193.240.157:9877 | ✓ 1014ms | 否 | ✓ 1054ms | ✓ 1973ms | 否 | http |
| 193.108.118.190:8888 | ✓ 1121ms | ✓ 1987ms | ✓ 1094ms | 否 | ✓ 1813ms | http |
| 103.104.99.89:80 | ✓ 1178ms | 否 | 否 | ✓ 1531ms | ✓ 1572ms | http |
| 103.111.136.82:34564 | 否 | 否 | ✓ 1865ms | ✓ 1592ms | ✓ 1604ms | http |
| 42.115.72.27:2039 | 否 | 否 | ✓ 1453ms | ✓ 1702ms | ✓ 1485ms | http |
| 185.191.236.162:3128 | ✓ 700ms | 否 | ✓ 688ms | ✓ 1758ms | ✓ 1449ms | http |
| 103.84.95.54:7890 | ✓ 1391ms | 否 | ✓ 1440ms | 否 | ✓ 1444ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1611ms | ✓ 991ms | 否 | ✓ 1280ms | http |
| 67.169.98.211:443 | ✓ 1332ms | 否 | ✓ 1433ms | 否 | ✓ 1858ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1315ms | 否 | ✓ 1433ms | ✓ 1122ms | http |
| 159.223.42.219:3128 | ✓ 866ms | 否 | ✓ 760ms | ✓ 1068ms | ✓ 868ms | http |
| 42.115.72.27:2102 | ✓ 1569ms | 否 | ✓ 1436ms | 否 | ✓ 1563ms | http |
| 14.225.222.164:7890 | ✓ 1670ms | ✓ 1756ms | ✓ 1378ms | ✓ 1936ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1230ms | ✓ 1290ms | ✓ 1304ms | ✓ 1166ms | ✓ 959ms | http |
| 116.80.82.226:3172 | ✓ 1934ms | 否 | 否 | ✓ 1864ms | ✓ 1873ms | http |
| 103.215.36.88:17633 | ✓ 1191ms | ✓ 1483ms | ✓ 1095ms | ✓ 1467ms | ✓ 1492ms | http |
| 121.230.9.252:1080 | 否 | 否 | ✓ 1794ms | ✓ 1781ms | ✓ 1230ms | http |
| 45.151.123.30:3128 | ✓ 1078ms | ✓ 1724ms | 否 | ✓ 1647ms | ✓ 1275ms | http |
| 69.48.179.20:3128 | ✓ 1432ms | 否 | ✓ 887ms | ✓ 1435ms | 否 | http |
| 180.103.19.143:1080 | 否 | 否 | ✓ 1027ms | ✓ 1554ms | ✓ 1156ms | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1796ms | ✓ 1559ms | ✓ 1156ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1866ms | ✓ 1953ms | ✓ 1704ms | http |
| 88.80.150.82:8080 | ✓ 1459ms | ✓ 1785ms | ✓ 894ms | 否 | 否 | https |
| 46.249.103.192:443 | ✓ 801ms | 否 | ✓ 1730ms | ✓ 1612ms | 否 | http |
| 172.105.212.216:8888 | ✓ 1707ms | 否 | ✓ 1475ms | ✓ 1204ms | ✓ 1455ms | http |
| 61.76.95.217:40088 | ✓ 1861ms | ✓ 1553ms | ✓ 1098ms | ✓ 1227ms | ✓ 1130ms | http |
| 211.171.114.154:3128 | ✓ 863ms | ✓ 1478ms | 否 | 否 | ✓ 1327ms | http |
| 121.230.9.54:1080 | ✓ 1061ms | ✓ 1423ms | ✓ 1300ms | ✓ 1443ms | ✓ 1177ms | http |
| 14.56.177.44:3128 | 否 | ✓ 1664ms | ✓ 1893ms | 否 | ✓ 1645ms | http |
| 176.100.39.18:3128 | ✓ 987ms | ✓ 1601ms | ✓ 1068ms | 否 | ✓ 1916ms | http |
| 120.92.212.16:7890 | ✓ 1223ms | ✓ 1205ms | ✓ 1008ms | 否 | ✓ 982ms | http |
| 116.80.48.16:7777 | ✓ 1584ms | 否 | ✓ 1661ms | ✓ 1832ms | 否 | http |
| 138.124.53.25:7443 | ✓ 728ms | 否 | ✓ 1119ms | 否 | ✓ 1636ms | http |
| 172.212.68.37:3128 | ✓ 947ms | ✓ 1778ms | ✓ 1422ms | ✓ 1977ms | ✓ 1141ms | http |
| 38.47.97.22:6005 | 否 | ✓ 890ms | ✓ 1511ms | ✓ 1085ms | ✓ 1019ms | http |
| 89.35.119.147:3128 | ✓ 812ms | ✓ 1760ms | ✓ 1499ms | 否 | ✓ 1693ms | http |
| 103.215.36.88:17406 | ✓ 1290ms | 否 | 否 | ✓ 1318ms | ✓ 1057ms | http |
| 103.139.138.194:3128 | ✓ 1089ms | 否 | ✓ 1141ms | ✓ 1533ms | ✓ 1629ms | http |
| 42.115.72.27:2046 | 否 | 否 | ✓ 1723ms | ✓ 1702ms | ✓ 1632ms | http |
| 173.212.222.244:8888 | ✓ 1062ms | ✓ 1524ms | ✓ 1560ms | 否 | 否 | http |
| 193.168.173.136:443 | ✓ 981ms | 否 | ✓ 1366ms | 否 | ✓ 1861ms | http |
| 42.115.72.27:2033 | 否 | 否 | ✓ 1626ms | ✓ 1709ms | ✓ 1871ms | http |
| 202.155.12.161:443 | ✓ 1808ms | 否 | ✓ 1329ms | ✓ 1157ms | 否 | http |

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
