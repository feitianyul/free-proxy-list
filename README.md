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

最后更新：2026-02-24 14:06:18 UTC（2026-02-24 22:06:18 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 472ms | ✓ 1195ms | 否 | ✓ 1652ms | ✓ 1018ms | http |
| 61.72.221.164:3128 | ✓ 1278ms | 否 | ✓ 1428ms | 否 | ✓ 1697ms | http |
| 124.16.93.70:7890 | ✓ 792ms | ✓ 1001ms | ✓ 822ms | ✓ 1044ms | ✓ 824ms | http |
| 34.50.41.78:8888 | ✓ 1536ms | 否 | ✓ 1521ms | ✓ 1520ms | ✓ 1184ms | http |
| 211.230.49.122:3128 | ✓ 899ms | ✓ 1629ms | ✓ 1676ms | ✓ 1490ms | ✓ 1492ms | http |
| 202.152.44.19:8081 | ✓ 1867ms | 否 | ✓ 1257ms | ✓ 1174ms | ✓ 1981ms | http |
| 202.152.44.18:8081 | ✓ 985ms | 否 | ✓ 952ms | ✓ 1292ms | ✓ 1036ms | http |
| 132.145.93.138:1080 | 否 | 否 | ✓ 1657ms | ✓ 1929ms | ✓ 1559ms | http |
| 120.92.212.16:8890 | ✓ 1115ms | ✓ 1184ms | ✓ 1135ms | ✓ 1177ms | ✓ 1718ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1154ms | ✓ 1434ms | ✓ 1123ms | http |
| 113.45.250.180:443 | ✓ 1091ms | 否 | ✓ 1048ms | ✓ 1115ms | ✓ 897ms | http |
| 20.210.76.178:8561 | ✓ 1946ms | ✓ 1968ms | ✓ 1598ms | 否 | 否 | http |
| 172.86.92.68:31337 | ✓ 806ms | 否 | ✓ 1327ms | ✓ 1875ms | 否 | http |
| 168.235.110.63:3128 | ✓ 437ms | ✓ 1214ms | ✓ 1336ms | ✓ 1305ms | ✓ 1015ms | http |
| 101.43.255.96:80 | ✓ 981ms | ✓ 1177ms | ✓ 1061ms | ✓ 1659ms | 否 | http |
| 81.70.169.194:80 | 否 | ✓ 1191ms | ✓ 1202ms | ✓ 1516ms | 否 | http |
| 35.225.22.61:80 | ✓ 1051ms | 否 | ✓ 1134ms | ✓ 1062ms | 否 | http |
| 20.27.15.49:8561 | ✓ 670ms | ✓ 1082ms | ✓ 705ms | ✓ 1028ms | ✓ 587ms | http |
| 20.210.76.175:8561 | ✓ 675ms | ✓ 1325ms | ✓ 660ms | ✓ 929ms | ✓ 729ms | http |
| 20.210.76.104:8561 | ✓ 670ms | ✓ 1583ms | ✓ 555ms | ✓ 835ms | ✓ 780ms | http |
| 5.101.0.233:3128 | ✓ 1414ms | 否 | ✓ 1867ms | 否 | ✓ 1837ms | http |
| 222.109.119.178:3128 | 否 | 否 | ✓ 806ms | ✓ 954ms | ✓ 1233ms | http |
| 36.136.27.2:4999 | ✓ 1276ms | ✓ 1747ms | ✓ 1694ms | 否 | 否 | http |
| 190.242.157.215:8080 | ✓ 1757ms | 否 | ✓ 900ms | ✓ 1936ms | 否 | http |
| 20.78.118.91:8561 | ✓ 1361ms | ✓ 1023ms | ✓ 591ms | ✓ 812ms | ✓ 725ms | http |
| 20.210.39.153:8561 | ✓ 1361ms | ✓ 1241ms | ✓ 470ms | ✓ 745ms | ✓ 702ms | http |
| 20.78.26.206:8561 | ✓ 1359ms | ✓ 1285ms | ✓ 507ms | ✓ 733ms | ✓ 660ms | http |
| 20.27.11.248:8561 | ✓ 1359ms | ✓ 1449ms | ✓ 582ms | ✓ 863ms | ✓ 649ms | http |
| 20.27.14.220:8561 | ✓ 1357ms | 否 | ✓ 435ms | ✓ 845ms | ✓ 643ms | http |
| 20.27.15.111:8561 | ✓ 1361ms | ✓ 1938ms | ✓ 437ms | ✓ 896ms | ✓ 652ms | http |
| 121.128.121.34:3128 | ✓ 739ms | ✓ 1254ms | ✓ 1188ms | ✓ 1451ms | ✓ 777ms | http |
| 178.253.22.108:65431 | ✓ 887ms | 否 | ✓ 1325ms | ✓ 1953ms | ✓ 1405ms | http |
| 183.98.143.134:8016 | ✓ 1389ms | 否 | 否 | ✓ 1133ms | ✓ 899ms | http |
| 121.128.121.44:3128 | ✓ 851ms | ✓ 1670ms | ✓ 1023ms | 否 | ✓ 952ms | http |
| 121.128.121.14:3128 | 否 | 否 | ✓ 632ms | ✓ 1532ms | ✓ 1921ms | http |
| 18.229.170.122:3128 | ✓ 1313ms | 否 | ✓ 752ms | 否 | ✓ 1670ms | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1743ms | ✓ 1402ms | ✓ 1328ms | http |
| 103.236.64.247:8888 | ✓ 1896ms | 否 | ✓ 1091ms | ✓ 1148ms | 否 | http |
| 18.229.201.117:3128 | ✓ 756ms | 否 | ✓ 846ms | 否 | ✓ 1749ms | http |
| 103.84.95.54:7890 | ✓ 1609ms | 否 | 否 | ✓ 1343ms | ✓ 640ms | http |
| 152.32.255.24:27197 | ✓ 1322ms | 否 | 否 | ✓ 1413ms | ✓ 1366ms | http |
| 120.46.152.136:3128 | ✓ 1065ms | ✓ 1322ms | ✓ 1394ms | ✓ 1747ms | ✓ 1064ms | http |
| 125.128.12.94:3128 | ✓ 641ms | 否 | ✓ 1125ms | ✓ 970ms | ✓ 1259ms | http |
| 160.238.65.5:3128 | ✓ 1023ms | ✓ 1817ms | 否 | 否 | ✓ 1545ms | http |
| 160.238.65.2:3128 | ✓ 1020ms | ✓ 1898ms | 否 | 否 | ✓ 1473ms | http |
| 160.238.65.9:3128 | ✓ 1011ms | ✓ 1888ms | 否 | 否 | ✓ 1486ms | http |
| 137.220.150.22:6005 | ✓ 820ms | 否 | ✓ 909ms | ✓ 1295ms | ✓ 957ms | http |
| 115.231.181.40:8128 | ✓ 842ms | 否 | 否 | ✓ 1869ms | ✓ 891ms | http |
| 101.32.244.83:8080 | ✓ 1002ms | 否 | ✓ 888ms | ✓ 1364ms | ✓ 1166ms | http |
| 121.43.196.210:8222 | ✓ 943ms | ✓ 1031ms | ✓ 813ms | ✓ 1048ms | ✓ 846ms | http |
| 121.43.196.213:8222 | ✓ 926ms | ✓ 991ms | ✓ 879ms | ✓ 1086ms | ✓ 816ms | http |
| 114.55.226.123:10086 | ✓ 994ms | ✓ 1555ms | ✓ 959ms | ✓ 1226ms | ✓ 980ms | http |
| 103.188.252.65:1234 | 否 | 否 | ✓ 1808ms | ✓ 1469ms | ✓ 1419ms | http |
| 217.216.109.116:8080 | 否 | 否 | ✓ 982ms | ✓ 1378ms | ✓ 1171ms | http |
| 101.47.73.135:3128 | ✓ 804ms | 否 | 否 | ✓ 1107ms | ✓ 1119ms | http |
| 35.234.17.221:8080 | ✓ 823ms | 否 | 否 | ✓ 1034ms | ✓ 1237ms | http |
| 180.127.149.225:1080 | ✓ 1934ms | ✓ 1172ms | ✓ 931ms | 否 | ✓ 1152ms | http |
| 85.208.108.43:2094 | 否 | 否 | ✓ 922ms | ✓ 1322ms | ✓ 1020ms | http |
| 62.113.119.14:8080 | ✓ 1405ms | 否 | ✓ 1655ms | 否 | ✓ 1960ms | http |
| 85.208.108.43:10808 | ✓ 507ms | 否 | ✓ 1041ms | ✓ 1414ms | ✓ 1070ms | http |
| 36.147.78.166:80 | ✓ 1655ms | 否 | ✓ 1717ms | 否 | ✓ 1613ms | http |
| 178.130.47.129:1082 | 否 | ✓ 1522ms | 否 | ✓ 799ms | ✓ 652ms | http |
| 217.216.109.116:80 | 否 | 否 | ✓ 1967ms | ✓ 1877ms | ✓ 1279ms | http |
| 217.76.245.80:999 | ✓ 852ms | 否 | ✓ 1847ms | ✓ 1469ms | ✓ 1487ms | http |
| 103.39.51.190:8080 | ✓ 1645ms | 否 | ✓ 1681ms | ✓ 1694ms | 否 | http |
| 121.128.121.74:3128 | ✓ 1755ms | 否 | 否 | ✓ 1987ms | ✓ 1248ms | http |
| 121.128.121.84:3128 | ✓ 1689ms | ✓ 1601ms | 否 | 否 | ✓ 1446ms | http |
| 36.147.78.166:443 | ✓ 1653ms | ✓ 1584ms | ✓ 1575ms | ✓ 1948ms | ✓ 1595ms | http |
| 14.56.177.182:3128 | 否 | 否 | ✓ 606ms | ✓ 1496ms | ✓ 1124ms | http |
| 14.56.107.244:3128 | ✓ 980ms | 否 | 否 | ✓ 1045ms | ✓ 778ms | http |
| 121.128.121.54:3128 | 否 | ✓ 1117ms | 否 | ✓ 1004ms | ✓ 744ms | http |
| 121.128.121.94:3128 | 否 | 否 | ✓ 1135ms | ✓ 1635ms | ✓ 1960ms | http |
| 59.127.212.110:4431 | 否 | ✓ 1106ms | 否 | ✓ 1081ms | ✓ 1910ms | http |
| 200.125.171.254:999 | ✓ 1478ms | 否 | ✓ 1560ms | ✓ 1509ms | ✓ 1339ms | http |
| 121.128.121.134:3128 | ✓ 1758ms | ✓ 1732ms | 否 | 否 | ✓ 983ms | http |
| 121.128.121.144:3128 | ✓ 1716ms | 否 | ✓ 1112ms | ✓ 1591ms | ✓ 977ms | http |

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
