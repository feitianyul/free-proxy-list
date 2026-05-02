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

最后更新：2026-05-02 04:16:34 UTC（2026-05-02 12:16:34 UTC+8）

**代理总数：87**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1924ms | 否 | ✓ 1437ms | ✓ 970ms | ✓ 767ms | http |
| 45.88.0.98:3128 | ✓ 1287ms | ✓ 1605ms | ✓ 1637ms | 否 | ✓ 1781ms | http |
| 45.88.0.117:3128 | ✓ 1292ms | ✓ 1722ms | ✓ 1517ms | 否 | ✓ 1799ms | http |
| 148.230.4.241:999 | ✓ 679ms | ✓ 1761ms | ✓ 487ms | 否 | ✓ 1396ms | http |
| 20.78.118.91:8561 | ✓ 445ms | ✓ 923ms | ✓ 572ms | ✓ 788ms | ✓ 642ms | http |
| 20.210.39.153:8561 | ✓ 450ms | ✓ 1011ms | ✓ 589ms | ✓ 757ms | ✓ 635ms | http |
| 20.78.26.206:8561 | ✓ 459ms | ✓ 1037ms | ✓ 564ms | ✓ 763ms | ✓ 626ms | http |
| 72.11.151.159:6005 | ✓ 667ms | ✓ 1496ms | ✓ 935ms | ✓ 1547ms | ✓ 1143ms | http |
| 213.220.62.62:3128 | ✓ 618ms | ✓ 1512ms | ✓ 897ms | ✓ 1568ms | ✓ 1257ms | http |
| 213.220.62.63:3128 | ✓ 652ms | 否 | ✓ 617ms | ✓ 1483ms | ✓ 1213ms | http |
| 103.35.190.69:1082 | ✓ 458ms | ✓ 1478ms | ✓ 926ms | 否 | 否 | http |
| 80.92.204.47:1081 | ✓ 608ms | ✓ 1488ms | ✓ 886ms | ✓ 1795ms | ✓ 1341ms | http |
| 218.108.131.186:17890 | ✓ 800ms | ✓ 1112ms | ✓ 835ms | ✓ 1167ms | ✓ 1952ms | http |
| 45.88.0.115:3128 | ✓ 599ms | ✓ 1557ms | ✓ 858ms | ✓ 1586ms | ✓ 1238ms | http |
| 34.96.238.40:8080 | ✓ 1123ms | 否 | ✓ 1383ms | 否 | ✓ 989ms | http |
| 45.88.0.116:3128 | ✓ 654ms | ✓ 1691ms | ✓ 681ms | ✓ 1515ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1651ms | 否 | ✓ 935ms | ✓ 1765ms | ✓ 937ms | http |
| 45.88.0.114:3128 | ✓ 621ms | 否 | ✓ 644ms | ✓ 1507ms | ✓ 1230ms | http |
| 113.160.132.26:8080 | ✓ 1976ms | ✓ 1373ms | ✓ 1281ms | ✓ 1634ms | ✓ 1418ms | http |
| 92.119.127.208:6005 | ✓ 1405ms | ✓ 1975ms | ✓ 1777ms | 否 | 否 | http |
| 212.58.132.5:8888 | ✓ 1604ms | 否 | ✓ 1691ms | ✓ 1743ms | ✓ 1310ms | http |
| 43.133.44.89:8888 | ✓ 1876ms | ✓ 1752ms | 否 | ✓ 1224ms | ✓ 969ms | http |
| 220.197.44.36:3128 | 否 | 否 | ✓ 1337ms | ✓ 1354ms | ✓ 1775ms | http |
| 45.167.124.71:999 | ✓ 745ms | ✓ 1759ms | ✓ 697ms | ✓ 1746ms | ✓ 1483ms | http |
| 47.85.51.197:1080 | ✓ 288ms | 否 | ✓ 313ms | ✓ 1295ms | ✓ 945ms | http |
| 45.88.0.99:3128 | ✓ 1208ms | ✓ 1550ms | ✓ 606ms | ✓ 1498ms | ✓ 1172ms | http |
| 45.88.0.113:3128 | ✓ 1215ms | ✓ 1646ms | ✓ 640ms | ✓ 1534ms | ✓ 1199ms | http |
| 45.88.0.111:3128 | ✓ 1208ms | 否 | ✓ 660ms | ✓ 1557ms | ✓ 1158ms | http |
| 20.127.128.70:8080 | ✓ 817ms | 否 | ✓ 870ms | ✓ 1684ms | ✓ 1301ms | http |
| 223.84.151.86:30005 | ✓ 1242ms | ✓ 1223ms | ✓ 1159ms | ✓ 1617ms | ✓ 1550ms | http |
| 120.92.108.86:7890 | ✓ 1362ms | 否 | ✓ 1794ms | 否 | ✓ 1359ms | http |
| 59.46.216.131:30001 | ✓ 1068ms | ✓ 1426ms | ✓ 1085ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 932ms | 否 | ✓ 958ms | 否 | ✓ 1242ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1621ms | ✓ 1699ms | ✓ 1460ms | ✓ 903ms | http |
| 34.246.183.20:9050 | ✓ 1409ms | 否 | ✓ 1931ms | 否 | ✓ 1527ms | http |
| 51.34.28.236:954 | ✓ 1412ms | 否 | ✓ 1923ms | 否 | ✓ 1553ms | http |
| 206.206.126.177:2412 | ✓ 1313ms | 否 | ✓ 857ms | ✓ 965ms | ✓ 965ms | http |
| 8.219.97.248:80 | ✓ 1483ms | 否 | ✓ 1248ms | ✓ 1667ms | 否 | http |
| 210.223.44.230:3128 | 否 | ✓ 1166ms | ✓ 968ms | ✓ 1008ms | ✓ 883ms | http |
| 91.184.241.12:443 | ✓ 913ms | 否 | ✓ 1614ms | 否 | ✓ 1964ms | http |
| 86.104.72.220:1082 | ✓ 313ms | ✓ 1373ms | ✓ 408ms | ✓ 1204ms | 否 | http |
| 46.101.95.183:8888 | ✓ 1697ms | ✓ 1959ms | ✓ 802ms | 否 | 否 | http |
| 45.140.147.82:1081 | ✓ 1680ms | ✓ 1681ms | ✓ 1266ms | ✓ 1981ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1326ms | 否 | ✓ 1194ms | 否 | ✓ 1529ms | http |
| 86.104.72.220:1081 | ✓ 376ms | ✓ 1328ms | ✓ 279ms | 否 | ✓ 860ms | http |
| 139.5.189.229:8888 | ✓ 1376ms | 否 | ✓ 1597ms | ✓ 1786ms | ✓ 1589ms | http |
| 23.236.65.144:8888 | ✓ 351ms | ✓ 608ms | ✓ 1220ms | ✓ 772ms | ✓ 547ms | http |
| 222.107.27.7:8062 | ✓ 1696ms | ✓ 1034ms | ✓ 946ms | ✓ 947ms | ✓ 738ms | http |
| 222.107.27.7:8053 | ✓ 1697ms | ✓ 1031ms | ✓ 946ms | ✓ 939ms | ✓ 757ms | http |
| 222.107.27.7:8089 | ✓ 1695ms | ✓ 1267ms | ✓ 712ms | ✓ 938ms | ✓ 949ms | http |
| 72.11.150.178:6005 | ✓ 429ms | ✓ 1283ms | ✓ 1089ms | 否 | 否 | http |
| 222.107.27.7:8092 | ✓ 1698ms | ✓ 1299ms | ✓ 1422ms | ✓ 1204ms | ✓ 741ms | http |
| 222.107.27.7:8085 | ✓ 1695ms | ✓ 1034ms | ✓ 947ms | ✓ 935ms | 否 | http |
| 222.107.27.7:8077 | ✓ 1698ms | ✓ 1009ms | ✓ 1080ms | ✓ 943ms | ✓ 749ms | http |
| 222.107.27.7:8088 | ✓ 1696ms | ✓ 1119ms | 否 | ✓ 966ms | 否 | http |
| 121.230.9.4:1080 | ✓ 1604ms | 否 | ✓ 970ms | ✓ 1392ms | ✓ 947ms | http |
| 121.230.9.11:1080 | ✓ 1064ms | ✓ 1586ms | ✓ 1199ms | 否 | 否 | http |
| 86.104.74.110:1081 | ✓ 1064ms | ✓ 1447ms | ✓ 1221ms | ✓ 1859ms | ✓ 1331ms | http |
| 63.179.134.206:40203 | ✓ 1417ms | 否 | ✓ 1980ms | 否 | ✓ 1982ms | http |
| 128.199.121.61:9090 | ✓ 1998ms | 否 | ✓ 968ms | ✓ 1311ms | 否 | http |
| 18.100.127.123:25053 | ✓ 1670ms | 否 | ✓ 1888ms | 否 | ✓ 1801ms | http |
| 8.154.21.175:3128 | ✓ 822ms | ✓ 1045ms | ✓ 838ms | ✓ 1100ms | ✓ 849ms | http |
| 101.32.243.189:80 | ✓ 1118ms | 否 | ✓ 1404ms | ✓ 1348ms | ✓ 1284ms | http |
| 103.119.229.140:8080 | ✓ 1667ms | 否 | ✓ 1545ms | ✓ 1781ms | ✓ 1795ms | http |
| 45.153.231.229:8080 | ✓ 1564ms | 否 | ✓ 1632ms | ✓ 1967ms | 否 | http |
| 183.238.3.150:7897 | ✓ 843ms | ✓ 1073ms | ✓ 919ms | ✓ 1053ms | 否 | http |
| 3.101.133.120:80 | ✓ 441ms | 否 | ✓ 1653ms | ✓ 1057ms | ✓ 641ms | http |
| 49.156.44.114:8080 | ✓ 1299ms | ✓ 1933ms | ✓ 1160ms | ✓ 1321ms | ✓ 1301ms | http |
| 15.160.116.45:8097 | ✓ 1792ms | 否 | ✓ 1679ms | 否 | ✓ 1923ms | http |
| 202.129.206.239:3128 | ✓ 1054ms | 否 | 否 | ✓ 1495ms | ✓ 1488ms | http |
| 150.107.140.238:3128 | ✓ 767ms | 否 | ✓ 892ms | ✓ 1471ms | ✓ 973ms | http |
| 86.104.72.219:1081 | ✓ 323ms | 否 | ✓ 304ms | ✓ 1872ms | 否 | http |
| 222.107.27.7:8074 | 否 | ✓ 1285ms | ✓ 1798ms | ✓ 1015ms | 否 | http |
| 92.119.127.211:6005 | ✓ 1322ms | 否 | ✓ 1664ms | 否 | ✓ 1588ms | http |
| 45.129.141.143:3128 | ✓ 1430ms | ✓ 1968ms | ✓ 1457ms | 否 | ✓ 1791ms | http |
| 223.16.170.103:3128 | ✓ 1130ms | ✓ 1819ms | ✓ 969ms | ✓ 1111ms | ✓ 1027ms | http |
| 130.61.174.200:1080 | ✓ 681ms | 否 | ✓ 1701ms | ✓ 1509ms | ✓ 1187ms | http |
| 43.167.237.94:3128 | ✓ 1212ms | ✓ 900ms | ✓ 504ms | ✓ 808ms | ✓ 620ms | http |
| 168.222.254.136:8888 | ✓ 1704ms | 否 | ✓ 1637ms | 否 | ✓ 1735ms | http |
| 45.140.147.155:1082 | ✓ 1588ms | 否 | ✓ 1680ms | 否 | ✓ 1832ms | http |
| 94.131.118.129:1081 | 否 | ✓ 1636ms | ✓ 570ms | 否 | ✓ 1527ms | http |
| 94.131.118.129:1082 | 否 | 否 | ✓ 567ms | ✓ 1924ms | ✓ 1211ms | http |
| 103.184.99.194:8080 | ✓ 1837ms | 否 | ✓ 1413ms | ✓ 1630ms | 否 | http |
| 61.52.131.172:8443 | ✓ 934ms | ✓ 1194ms | ✓ 918ms | ✓ 1164ms | ✓ 895ms | http |
| 54.229.201.146:9091 | ✓ 966ms | 否 | ✓ 1585ms | 否 | ✓ 1859ms | http |
| 15.160.132.166:24608 | ✓ 1936ms | 否 | ✓ 872ms | ✓ 1956ms | ✓ 1624ms | http |
| 103.39.51.207:8080 | ✓ 1370ms | 否 | 否 | ✓ 1712ms | ✓ 1478ms | http |

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
