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

最后更新：2026-05-03 15:41:58 UTC（2026-05-03 23:41:58 UTC+8）

**代理总数：70**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 841ms | ✓ 1693ms | ✓ 922ms | ✓ 1089ms | ✓ 1941ms | http |
| 80.92.204.47:1081 | ✓ 790ms | ✓ 1723ms | ✓ 987ms | ✓ 1684ms | ✓ 1224ms | http |
| 1.231.81.166:3128 | ✓ 1952ms | ✓ 1825ms | ✓ 1774ms | ✓ 1314ms | ✓ 1272ms | http |
| 148.230.4.241:999 | ✓ 699ms | ✓ 1994ms | ✓ 816ms | ✓ 1656ms | 否 | http |
| 20.27.11.248:8561 | ✓ 648ms | ✓ 1103ms | ✓ 718ms | ✓ 1042ms | ✓ 799ms | http |
| 20.27.14.220:8561 | ✓ 665ms | ✓ 1277ms | ✓ 707ms | ✓ 1037ms | ✓ 801ms | http |
| 154.64.232.35:8080 | ✓ 374ms | ✓ 800ms | 否 | ✓ 1582ms | ✓ 811ms | http |
| 109.120.156.122:8090 | ✓ 1184ms | 否 | ✓ 779ms | 否 | ✓ 1641ms | http |
| 152.32.132.190:7890 | ✓ 1297ms | ✓ 1219ms | 否 | ✓ 1126ms | ✓ 1044ms | http |
| 150.249.255.91:3128 | ✓ 704ms | 否 | 否 | ✓ 1036ms | ✓ 821ms | http |
| 113.160.132.26:8080 | ✓ 1721ms | ✓ 1629ms | ✓ 1555ms | ✓ 1740ms | ✓ 1575ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1799ms | ✓ 1554ms | ✓ 1247ms | http |
| 46.105.190.38:3128 | ✓ 1422ms | 否 | ✓ 390ms | ✓ 1720ms | 否 | http |
| 45.167.124.71:999 | ✓ 1336ms | 否 | ✓ 438ms | ✓ 1585ms | ✓ 1348ms | http |
| 193.123.250.39:1080 | 否 | 否 | ✓ 1531ms | ✓ 1830ms | ✓ 982ms | http |
| 190.12.150.244:999 | ✓ 1386ms | 否 | ✓ 1827ms | ✓ 1913ms | ✓ 1546ms | http |
| 20.78.26.206:8561 | ✓ 652ms | ✓ 979ms | ✓ 718ms | ✓ 1080ms | ✓ 819ms | http |
| 20.78.118.91:8561 | ✓ 652ms | ✓ 1214ms | ✓ 697ms | ✓ 1009ms | ✓ 815ms | http |
| 20.210.39.153:8561 | ✓ 677ms | ✓ 1318ms | ✓ 717ms | ✓ 981ms | ✓ 804ms | http |
| 20.27.13.35:8561 | ✓ 983ms | ✓ 1171ms | ✓ 637ms | ✓ 1047ms | ✓ 897ms | http |
| 20.210.76.175:8561 | ✓ 781ms | ✓ 1229ms | ✓ 705ms | ✓ 1004ms | ✓ 791ms | http |
| 20.18.193.135:8561 | ✓ 778ms | ✓ 1245ms | ✓ 672ms | ✓ 1022ms | ✓ 815ms | http |
| 20.210.76.104:8561 | ✓ 779ms | ✓ 1064ms | ✓ 781ms | ✓ 1131ms | ✓ 889ms | http |
| 20.210.76.178:8561 | ✓ 831ms | ✓ 1371ms | ✓ 680ms | ✓ 1123ms | ✓ 820ms | http |
| 20.27.15.111:8561 | ✓ 982ms | ✓ 1442ms | ✓ 876ms | ✓ 1126ms | ✓ 804ms | http |
| 20.27.15.49:8561 | ✓ 765ms | 否 | ✓ 686ms | ✓ 1010ms | ✓ 826ms | http |
| 62.113.119.14:8080 | ✓ 1273ms | ✓ 1435ms | ✓ 1006ms | 否 | 否 | http |
| 206.206.126.177:2412 | ✓ 1703ms | 否 | 否 | ✓ 1373ms | ✓ 1029ms | http |
| 91.233.223.147:3128 | ✓ 1661ms | ✓ 1901ms | ✓ 1801ms | 否 | ✓ 1915ms | http |
| 46.105.190.40:3128 | ✓ 416ms | ✓ 1224ms | ✓ 655ms | ✓ 1673ms | ✓ 1270ms | http |
| 37.187.109.70:10111 | ✓ 1045ms | 否 | ✓ 819ms | ✓ 1781ms | ✓ 1351ms | http |
| 45.140.147.82:1082 | ✓ 1248ms | 否 | ✓ 1019ms | 否 | ✓ 1898ms | http |
| 185.21.11.140:1080 | ✓ 845ms | 否 | ✓ 1831ms | 否 | ✓ 1796ms | http |
| 45.78.79.225:1080 | ✓ 1497ms | 否 | ✓ 1256ms | 否 | ✓ 1884ms | http |
| 103.157.200.126:3128 | ✓ 1203ms | 否 | ✓ 1241ms | 否 | ✓ 1831ms | http |
| 147.45.178.211:14658 | ✓ 1666ms | 否 | ✓ 935ms | ✓ 1390ms | 否 | http |
| 38.180.2.107:3128 | ✓ 945ms | 否 | ✓ 1810ms | 否 | ✓ 1743ms | http |
| 41.33.219.130:1981 | ✓ 1398ms | 否 | ✓ 1638ms | 否 | ✓ 1670ms | http |
| 104.247.51.76:3128 | 否 | 否 | ✓ 784ms | ✓ 1031ms | ✓ 786ms | http |
| 178.156.224.42:3128 | ✓ 1890ms | 否 | ✓ 1631ms | 否 | ✓ 1511ms | http |
| 77.110.107.80:8080 | ✓ 1355ms | ✓ 1994ms | ✓ 1297ms | ✓ 1707ms | 否 | http |
| 77.110.107.80:1080 | ✓ 1394ms | 否 | ✓ 1290ms | ✓ 1708ms | 否 | http |
| 8.154.21.175:3128 | ✓ 1907ms | ✓ 1041ms | ✓ 947ms | ✓ 1180ms | ✓ 944ms | http |
| 103.165.138.173:8181 | 否 | 否 | ✓ 1336ms | ✓ 1402ms | ✓ 1116ms | http |
| 130.61.174.200:1080 | ✓ 1022ms | 否 | ✓ 1667ms | ✓ 1302ms | 否 | http |
| 8.219.97.248:80 | ✓ 1623ms | 否 | ✓ 1624ms | 否 | ✓ 1748ms | http |
| 45.125.67.37:8443 | ✓ 1074ms | 否 | ✓ 1099ms | ✓ 1376ms | ✓ 1355ms | http |
| 212.58.132.5:8888 | ✓ 1481ms | 否 | ✓ 1623ms | ✓ 1587ms | ✓ 1266ms | http |
| 47.77.216.82:1080 | ✓ 309ms | 否 | ✓ 726ms | 否 | ✓ 820ms | http |
| 3.101.133.120:80 | ✓ 554ms | ✓ 1534ms | ✓ 957ms | 否 | ✓ 1947ms | http |
| 20.127.128.70:8080 | ✓ 1454ms | 否 | ✓ 414ms | ✓ 1161ms | ✓ 962ms | http |
| 142.93.195.158:80 | ✓ 165ms | ✓ 1395ms | 否 | ✓ 1466ms | ✓ 1717ms | http |
| 20.205.16.149:3128 | ✓ 876ms | ✓ 1626ms | ✓ 1004ms | ✓ 1358ms | ✓ 953ms | http |
| 20.2.83.243:3128 | ✓ 834ms | 否 | 否 | ✓ 1258ms | ✓ 910ms | http |
| 134.209.153.66:3128 | ✓ 1680ms | 否 | ✓ 1651ms | ✓ 1917ms | ✓ 1608ms | http |
| 94.131.118.129:1081 | ✓ 551ms | ✓ 1994ms | ✓ 879ms | 否 | 否 | http |
| 43.133.44.89:8888 | ✓ 1721ms | 否 | ✓ 1136ms | ✓ 1266ms | ✓ 1015ms | http |
| 121.147.253.205:3124 | ✓ 1633ms | 否 | ✓ 1798ms | 否 | ✓ 1565ms | http |
| 103.209.36.58:8080 | ✓ 1400ms | 否 | 否 | ✓ 1718ms | ✓ 1930ms | http |
| 62.60.149.161:3128 | ✓ 1017ms | 否 | ✓ 1241ms | ✓ 1944ms | ✓ 1489ms | http |
| 42.200.76.16:3888 | ✓ 942ms | 否 | ✓ 870ms | ✓ 1310ms | ✓ 841ms | http |
| 86.104.74.110:1081 | ✓ 1356ms | ✓ 1292ms | ✓ 368ms | ✓ 1394ms | ✓ 1056ms | http |
| 138.124.108.176:3128 | ✓ 1617ms | ✓ 1453ms | ✓ 1384ms | ✓ 1680ms | ✓ 1414ms | http |
| 104.248.195.47:8080 | ✓ 411ms | 否 | ✓ 1984ms | ✓ 1405ms | ✓ 1337ms | http |
| 61.52.131.172:8443 | ✓ 925ms | 否 | ✓ 1004ms | ✓ 1521ms | ✓ 1100ms | http |
| 34.96.238.40:8080 | ✓ 1720ms | ✓ 1407ms | 否 | 否 | ✓ 1309ms | http |
| 94.131.118.129:1082 | ✓ 793ms | ✓ 1502ms | ✓ 518ms | ✓ 1354ms | ✓ 864ms | http |
| 220.197.44.36:3128 | ✓ 1244ms | ✓ 1337ms | ✓ 1326ms | ✓ 1541ms | ✓ 1288ms | http |
| 173.212.245.136:8888 | ✓ 1181ms | ✓ 1593ms | 否 | 否 | ✓ 1601ms | http |
| 144.31.25.69:21064 | ✓ 1335ms | 否 | ✓ 869ms | 否 | ✓ 1495ms | http |

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
