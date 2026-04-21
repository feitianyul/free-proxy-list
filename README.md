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

最后更新：2026-04-21 14:43:28 UTC（2026-04-21 22:43:28 UTC+8）

**代理总数：67**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 1477ms | ✓ 1355ms | 否 | ✓ 1148ms | ✓ 909ms | http |
| 1.231.81.166:3128 | ✓ 1519ms | ✓ 1113ms | ✓ 865ms | ✓ 907ms | ✓ 730ms | http |
| 149.51.42.10:3128 | ✓ 1526ms | ✓ 1887ms | 否 | ✓ 1836ms | 否 | http |
| 152.42.208.139:8118 | ✓ 1527ms | 否 | ✓ 1249ms | ✓ 1118ms | ✓ 888ms | http |
| 46.101.95.183:8888 | ✓ 758ms | 否 | ✓ 1546ms | 否 | ✓ 1433ms | http |
| 113.160.132.26:8080 | ✓ 1849ms | ✓ 1369ms | ✓ 1213ms | ✓ 1387ms | ✓ 1025ms | http |
| 162.19.253.202:8443 | ✓ 991ms | 否 | ✓ 1928ms | 否 | ✓ 1998ms | http |
| 149.51.42.10:8080 | ✓ 1532ms | ✓ 1617ms | 否 | ✓ 1468ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1292ms | 否 | ✓ 1569ms | ✓ 1434ms | 否 | http |
| 177.93.132.244:3128 | ✓ 781ms | 否 | ✓ 767ms | 否 | ✓ 1805ms | http |
| 188.246.224.49:7890 | ✓ 774ms | ✓ 1617ms | ✓ 1907ms | ✓ 1865ms | ✓ 1778ms | http |
| 212.58.132.5:8888 | ✓ 1137ms | 否 | ✓ 1518ms | ✓ 1474ms | ✓ 1214ms | http |
| 81.30.156.115:8080 | ✓ 1044ms | 否 | ✓ 1050ms | 否 | ✓ 1405ms | http |
| 14.143.222.113:57788 | 否 | 否 | ✓ 1165ms | ✓ 1358ms | ✓ 1941ms | http |
| 20.127.128.70:8080 | ✓ 936ms | 否 | ✓ 803ms | ✓ 1595ms | ✓ 1458ms | http |
| 91.99.15.45:2095 | ✓ 769ms | 否 | ✓ 1763ms | 否 | ✓ 1587ms | http |
| 185.138.116.150:8080 | ✓ 1549ms | 否 | ✓ 1296ms | ✓ 1994ms | 否 | http |
| 152.70.91.193:40000 | ✓ 1405ms | 否 | 否 | ✓ 1677ms | ✓ 1523ms | http |
| 20.78.118.91:8561 | ✓ 512ms | ✓ 1302ms | ✓ 721ms | ✓ 932ms | ✓ 830ms | http |
| 8.219.195.129:1080 | ✓ 817ms | 否 | ✓ 796ms | ✓ 1088ms | ✓ 852ms | http |
| 146.190.80.158:9090 | ✓ 1030ms | 否 | ✓ 1293ms | ✓ 1316ms | ✓ 1140ms | http |
| 45.76.207.177:40000 | ✓ 1158ms | 否 | ✓ 932ms | ✓ 1838ms | ✓ 1477ms | http |
| 128.199.116.219:9090 | ✓ 1848ms | 否 | ✓ 785ms | ✓ 1439ms | ✓ 973ms | http |
| 168.144.75.9:3128 | ✓ 1163ms | 否 | ✓ 1518ms | 否 | ✓ 1991ms | http |
| 20.78.26.206:8561 | ✓ 512ms | ✓ 1307ms | ✓ 537ms | ✓ 925ms | ✓ 727ms | http |
| 34.71.229.255:3128 | ✓ 442ms | 否 | ✓ 970ms | ✓ 960ms | ✓ 750ms | http |
| 20.210.39.153:8561 | ✓ 521ms | ✓ 1231ms | ✓ 556ms | ✓ 843ms | ✓ 695ms | http |
| 159.89.191.221:3128 | ✓ 1416ms | 否 | ✓ 652ms | ✓ 1164ms | ✓ 864ms | http |
| 210.77.29.244:6478 | ✓ 955ms | ✓ 1841ms | ✓ 971ms | ✓ 1140ms | ✓ 942ms | http |
| 160.238.65.6:3128 | ✓ 1719ms | ✓ 1439ms | ✓ 1259ms | 否 | ✓ 1855ms | http |
| 84.47.150.125:1080 | ✓ 1492ms | 否 | 否 | ✓ 1765ms | ✓ 1817ms | http |
| 160.238.65.2:3128 | ✓ 1714ms | 否 | ✓ 685ms | 否 | ✓ 1874ms | http |
| 210.223.44.230:3128 | ✓ 1658ms | 否 | 否 | ✓ 955ms | ✓ 747ms | http |
| 45.153.231.229:8080 | ✓ 1169ms | 否 | ✓ 1867ms | 否 | ✓ 1733ms | http |
| 62.113.119.14:8080 | ✓ 1441ms | 否 | ✓ 1451ms | 否 | ✓ 1507ms | http |
| 188.253.125.38:28798 | ✓ 981ms | ✓ 1302ms | ✓ 1160ms | ✓ 1285ms | ✓ 881ms | http |
| 117.236.124.166:3128 | ✓ 1052ms | 否 | ✓ 1091ms | ✓ 1919ms | 否 | http |
| 218.108.131.186:17890 | 否 | ✓ 1295ms | ✓ 1478ms | ✓ 1561ms | ✓ 1167ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1558ms | 否 | ✓ 963ms | ✓ 1215ms | http |
| 60.249.94.208:3128 | ✓ 765ms | ✓ 1067ms | ✓ 866ms | ✓ 973ms | ✓ 740ms | http |
| 124.16.103.90:10808 | ✓ 1105ms | ✓ 1196ms | ✓ 1082ms | ✓ 1204ms | ✓ 991ms | http |
| 159.203.220.84:3128 | 否 | 否 | ✓ 911ms | ✓ 1124ms | ✓ 789ms | http |
| 118.193.44.243:3128 | ✓ 1109ms | 否 | ✓ 923ms | 否 | ✓ 714ms | http |
| 114.237.77.219:1080 | ✓ 1015ms | 否 | ✓ 1432ms | ✓ 1584ms | ✓ 1042ms | http |
| 223.84.151.86:30005 | ✓ 1313ms | ✓ 1427ms | ✓ 1399ms | 否 | 否 | http |
| 160.238.65.7:3128 | ✓ 651ms | ✓ 1536ms | 否 | 否 | ✓ 1762ms | http |
| 84.47.150.126:1080 | ✓ 1400ms | 否 | 否 | ✓ 1916ms | ✓ 1395ms | http |
| 89.208.106.138:10808 | ✓ 1267ms | ✓ 1659ms | ✓ 889ms | 否 | 否 | http |
| 116.171.106.26:3443 | ✓ 1722ms | ✓ 1647ms | ✓ 1909ms | 否 | 否 | http |
| 208.87.243.199:7878 | 否 | ✓ 1002ms | ✓ 1389ms | ✓ 1466ms | 否 | http |
| 42.101.8.101:8888 | 否 | 否 | ✓ 1152ms | ✓ 1838ms | ✓ 1147ms | http |
| 20.210.76.175:8561 | ✓ 555ms | ✓ 1034ms | ✓ 855ms | ✓ 937ms | ✓ 662ms | http |
| 20.18.193.135:8561 | ✓ 556ms | ✓ 1157ms | ✓ 732ms | ✓ 935ms | ✓ 661ms | http |
| 20.27.15.49:8561 | ✓ 560ms | ✓ 1661ms | ✓ 493ms | ✓ 951ms | ✓ 763ms | http |
| 20.210.76.178:8561 | ✓ 531ms | ✓ 1709ms | ✓ 499ms | ✓ 928ms | ✓ 813ms | http |
| 43.132.188.134:443 | ✓ 687ms | 否 | 否 | ✓ 1491ms | ✓ 904ms | http |
| 47.84.73.61:1080 | ✓ 1478ms | 否 | ✓ 754ms | ✓ 1188ms | ✓ 876ms | http |
| 20.27.11.248:8561 | ✓ 1424ms | ✓ 1321ms | ✓ 1547ms | ✓ 1137ms | ✓ 838ms | http |
| 20.27.15.111:8561 | ✓ 1425ms | 否 | ✓ 1051ms | ✓ 1036ms | ✓ 838ms | http |
| 115.231.181.40:8128 | ✓ 1868ms | ✓ 1835ms | ✓ 1390ms | 否 | 否 | http |
| 20.27.13.35:8561 | ✓ 604ms | 否 | ✓ 636ms | ✓ 812ms | ✓ 805ms | http |
| 103.113.70.189:1081 | ✓ 630ms | 否 | ✓ 376ms | ✓ 1785ms | ✓ 1384ms | http |
| 59.46.216.131:30001 | ✓ 1459ms | ✓ 1353ms | ✓ 1403ms | 否 | 否 | http |
| 20.27.14.220:8561 | ✓ 551ms | 否 | ✓ 762ms | ✓ 1058ms | ✓ 827ms | http |
| 103.69.84.106:3131 | 否 | 否 | ✓ 1450ms | ✓ 1170ms | ✓ 1149ms | http |
| 111.79.111.126:3128 | ✓ 1231ms | ✓ 1365ms | ✓ 1597ms | 否 | 否 | http |
| 103.56.115.156:7890 | ✓ 1543ms | 否 | 否 | ✓ 1523ms | ✓ 1482ms | http |

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
