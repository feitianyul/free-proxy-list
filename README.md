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

最后更新：2026-03-16 17:57:35 UTC（2026-03-17 01:57:35 UTC+8）

**代理总数：53**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 5.129.206.247:8888 | ✓ 1386ms | ✓ 1719ms | ✓ 1520ms | 否 | 否 | http |
| 147.161.210.140:8800 | ✓ 1648ms | ✓ 1877ms | ✓ 1076ms | ✓ 1707ms | ✓ 1085ms | http |
| 1.231.81.166:3128 | ✓ 1684ms | ✓ 1294ms | ✓ 1751ms | ✓ 1538ms | ✓ 952ms | http |
| 45.136.198.40:3128 | 否 | ✓ 1817ms | ✓ 1552ms | 否 | ✓ 1980ms | http |
| 113.160.132.26:8080 | ✓ 1954ms | 否 | ✓ 1474ms | ✓ 1528ms | 否 | http |
| 103.113.70.189:1081 | ✓ 879ms | 否 | 否 | ✓ 1181ms | ✓ 729ms | http |
| 137.220.150.104:6005 | ✓ 1767ms | 否 | ✓ 1746ms | ✓ 1791ms | ✓ 1448ms | http |
| 202.155.12.161:443 | ✓ 991ms | 否 | ✓ 1218ms | ✓ 1172ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1576ms | ✓ 1864ms | ✓ 1242ms | ✓ 1529ms | ✓ 1305ms | http |
| 101.43.127.100:8877 | ✓ 1865ms | ✓ 1303ms | ✓ 998ms | ✓ 1262ms | ✓ 1092ms | http |
| 62.60.177.204:34094 | ✓ 730ms | ✓ 1241ms | 否 | ✓ 1235ms | ✓ 722ms | http |
| 8.209.239.31:30000 | ✓ 646ms | ✓ 1270ms | 否 | ✓ 944ms | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1255ms | 否 | ✓ 1130ms | ✓ 1084ms | http |
| 137.220.150.152:6005 | ✓ 1776ms | 否 | ✓ 1062ms | ✓ 1504ms | ✓ 1264ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1708ms | ✓ 1293ms | 否 | ✓ 1223ms | http |
| 185.115.74.185:8080 | ✓ 930ms | ✓ 1773ms | ✓ 1948ms | 否 | 否 | http |
| 72.11.151.159:6005 | ✓ 500ms | 否 | ✓ 821ms | ✓ 1198ms | ✓ 808ms | http |
| 210.77.29.244:6789 | ✓ 1003ms | ✓ 1348ms | ✓ 1096ms | ✓ 1300ms | ✓ 1033ms | http |
| 137.220.151.110:6005 | ✓ 1484ms | 否 | ✓ 1291ms | ✓ 1300ms | ✓ 1544ms | http |
| 92.119.127.213:6005 | ✓ 1070ms | 否 | ✓ 1657ms | 否 | ✓ 1512ms | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 1187ms | ✓ 1432ms | ✓ 1034ms | http |
| 212.192.12.90:6005 | 否 | 否 | ✓ 1507ms | ✓ 1927ms | ✓ 1149ms | http |
| 105.159.133.152:4133 | ✓ 1542ms | 否 | ✓ 1910ms | 否 | ✓ 1656ms | http |
| 86.53.183.16:1080 | ✓ 513ms | ✓ 1697ms | ✓ 1388ms | 否 | 否 | http |
| 85.198.96.242:3128 | ✓ 717ms | ✓ 1994ms | ✓ 936ms | 否 | 否 | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1585ms | ✓ 1845ms | ✓ 1421ms | http |
| 149.50.116.240:1080 | ✓ 1983ms | ✓ 1725ms | ✓ 1864ms | 否 | ✓ 1360ms | http |
| 103.82.23.118:5171 | ✓ 1537ms | 否 | ✓ 1794ms | 否 | ✓ 1659ms | http |
| 219.117.204.211:7799 | ✓ 1514ms | 否 | ✓ 1140ms | 否 | ✓ 1162ms | http |
| 210.223.44.230:3128 | ✓ 1628ms | 否 | ✓ 1112ms | ✓ 1118ms | ✓ 910ms | http |
| 88.80.150.82:8080 | ✓ 1130ms | 否 | ✓ 1865ms | 否 | ✓ 1499ms | https |
| 115.231.181.40:8128 | 否 | ✓ 1378ms | ✓ 1263ms | ✓ 1444ms | ✓ 1074ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1432ms | ✓ 1191ms | 否 | ✓ 1176ms | http |
| 106.117.208.101:7890 | ✓ 1172ms | ✓ 1512ms | ✓ 1222ms | ✓ 1471ms | ✓ 1187ms | http |
| 38.34.179.78:8448 | 否 | 否 | ✓ 443ms | ✓ 1077ms | ✓ 717ms | http |
| 120.92.212.16:8890 | ✓ 1137ms | ✓ 1405ms | ✓ 1110ms | 否 | ✓ 1172ms | http |
| 38.34.179.60:8450 | ✓ 1099ms | ✓ 1108ms | ✓ 773ms | ✓ 1404ms | ✓ 1154ms | http |
| 38.34.179.14:8450 | ✓ 626ms | ✓ 1529ms | ✓ 532ms | ✓ 1164ms | ✓ 834ms | http |
| 186.148.180.46:999 | ✓ 1036ms | ✓ 1789ms | ✓ 1515ms | ✓ 1875ms | ✓ 1548ms | http |
| 14.225.212.37:7890 | ✓ 1158ms | ✓ 1952ms | 否 | 否 | ✓ 1876ms | http |
| 8.219.97.248:80 | ✓ 1927ms | 否 | ✓ 1748ms | 否 | ✓ 1490ms | http |
| 38.34.179.58:8450 | ✓ 667ms | ✓ 921ms | ✓ 1288ms | ✓ 991ms | ✓ 879ms | http |
| 137.184.6.37:3128 | 否 | ✓ 1469ms | ✓ 1150ms | ✓ 1443ms | 否 | http |
| 133.242.138.34:8100 | ✓ 1450ms | 否 | ✓ 1913ms | 否 | ✓ 1593ms | http |
| 168.235.110.63:3128 | ✓ 747ms | ✓ 1028ms | ✓ 681ms | ✓ 1440ms | ✓ 1547ms | http |
| 198.24.188.139:31029 | ✓ 535ms | ✓ 1794ms | ✓ 1465ms | ✓ 1462ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1659ms | ✓ 1506ms | 否 | ✓ 1696ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1986ms | 否 | 否 | ✓ 1560ms | ✓ 1605ms | http |
| 178.156.224.42:3128 | ✓ 1526ms | 否 | ✓ 1463ms | 否 | ✓ 1902ms | http |
| 116.6.106.33:3128 | ✓ 894ms | ✓ 1240ms | ✓ 956ms | ✓ 1182ms | ✓ 925ms | http |
| 121.230.9.231:1080 | ✓ 1457ms | ✓ 1645ms | ✓ 1439ms | ✓ 1642ms | ✓ 1195ms | http |
| 45.136.130.176:8451 | 否 | ✓ 934ms | ✓ 1255ms | ✓ 952ms | ✓ 712ms | http |
| 2.56.122.146:10808 | ✓ 933ms | ✓ 1683ms | ✓ 1457ms | ✓ 1783ms | ✓ 1214ms | http |

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
